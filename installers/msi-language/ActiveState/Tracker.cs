﻿using GAPixel;
using GoogleAnalyticsTracker.Core;
using GoogleAnalyticsTracker.Core.TrackerParameters;
using GoogleAnalyticsTracker.Simple;
using Microsoft.Deployment.WindowsInstaller;
using System;
using System.Net;
using System.Threading.Tasks;


namespace ActiveState
{
    public sealed class TrackerSingleton
    {
        private static readonly Lazy<TrackerSingleton> lazy = new Lazy<TrackerSingleton>(() => new TrackerSingleton());
        private static string GoogleAnalyticsUserAgent = "UA-118120158-2";

        private readonly SimpleTracker _tracker;
        private readonly string _cid;

        public static TrackerSingleton Instance { get { return lazy.Value; } }

        public TrackerSingleton()
        {
            var simpleTrackerEnvironment = new SimpleTrackerEnvironment(Environment.OSVersion.Platform.ToString(),
                Environment.OSVersion.Version.ToString(),
                Environment.OSVersion.VersionString);
            this._tracker = new SimpleTracker(GoogleAnalyticsUserAgent, simpleTrackerEnvironment);
            this._cid = GetInfo.GetUniqueId();
        }

        public async Task<TrackingResult> TrackEventAsync(string sessionID, string category, string action, string label, string msiVersion, long value = 1)
        {
            var eventTrackingParameters = new EventTracking
            {
                Category = category,
                Action = action,
                Label = label,
                Value = value,
            };

            eventTrackingParameters.ClientId = this._cid;
            eventTrackingParameters.SetCustomDimensions(new System.Collections.Generic.Dictionary<int, string> {
                { 1, msiVersion },
                { 2, sessionID },
            });

            return await this._tracker.TrackAsync(eventTrackingParameters);
        }

        public async Task TrackS3Event(Session session, string sessionID, string category, string action, string label)
        {
            string pixelURL = string.Format(
                "https://cli-msi.s3.amazonaws.com/pixel.txt?x-referrer={0}&x-session={1}&x-event={2}&x-event-category={3}&x-event-value={4}",
                this._cid, sessionID, action, category, label
            );
            session.Log(string.Format("Downloading S3 pixel from URL: {0}", pixelURL));
            try
            {
                WebClient client = new WebClient();
                await client.DownloadStringTaskAsync(pixelURL);
            }
            catch (WebException e)
            {
                string msg = string.Format("Encountered exception downloading S3 pixel file: {0}", e.ToString());
                session.Log(msg);
                RollbarReport.Error(msg, session);
            }

            session.Log("Successfully downloaded S3 pixel string");
        }

        /// <summary>
        /// Sends a GA event in background (fires and forgets)
        /// </summary>
        /// <description>
        /// The event can fail to be send if the main process gets cancelled before the task finishes.
        /// Use the synchronous version of this command in that case.
        /// </description>
        public void TrackEventInBackground(Session session, string sessionID, string category, string action, string label, string langVersion, long value = 1)
        {
            var pid = System.Diagnostics.Process.GetCurrentProcess().Id;

            session.Log("Sending background event {0}/{1}/{2} for cid={3} (custom dimension 1: {4}, pid={5})", category, action, label, this._cid, langVersion, pid);
            Task.Run(async () =>
            {
                await TrackEventAsync(sessionID, category, action, label, langVersion, value);
                await TrackS3Event(session, sessionID, category, action, label);
            });
        }

        /// <summary>
        /// Sends a GA event and waits for the request to complete.
        /// </summary>
        public void TrackEventSynchronously(Session session, string sessionID, string category, string action, string label, string langVersion, long value = 1)
        {
            var pid = System.Diagnostics.Process.GetCurrentProcess().Id;

            session.Log("Sending event {0}/{1}/{2} for cid={3} (custom dimension 1: {4}, pid={5})", category, action, label, this._cid, langVersion, pid);
            var t = Task.Run(async () =>
            {
                await TrackEventAsync(sessionID, category, action, label, langVersion, value);
                await TrackS3Event(session, sessionID, category, action, label);
            });
            t.Wait();
        }
    }
};
