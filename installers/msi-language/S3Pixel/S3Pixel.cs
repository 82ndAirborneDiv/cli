using GAPixel;
using System.Net;
using Microsoft.Deployment.WindowsInstaller;
using ActiveState;

namespace S3Pixel
{
    public class CustomActions
    {
        [CustomAction]
        public static ActionResult DownloadPixel(Session session)
        {
            RollbarHelper.ConfigureRollbarSingleton(session["COMMIT_ID"]);

            using (var log = new Logging(session, session["INSTALLDIR"]))
            {
                log.Log("Begin download S3 pixel");

                string guid = GetInfo.GetUniqueId(session);
                string pixelURL = string.Format("https://cli-msi.s3.amazonaws.com/pixel.txt?x-referrer={0}", guid);
                log.Log(string.Format("Downloading S3 pixel from URL: {0}", pixelURL));
                try
                {
                    WebClient client = new WebClient();
                    client.DownloadString(pixelURL);
                }
                catch (WebException e)
                {
                    string msg = string.Format("Encountered exception downloading S3 pixel file: {0}", e.ToString());
                    log.Log(msg);
                    RollbarReport.Error(msg, log);
                }

                log.Log("Successfully downloaded S3 pixel string");
                return ActionResult.Success;
            }
        }
    }
}
