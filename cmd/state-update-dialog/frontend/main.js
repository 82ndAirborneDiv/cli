// Main entry point
function start() {
    showContextMenu = document.oncontextmenu;
    document.oncontextmenu = () => false;
    backend.Bindings.DebugMode().then(debugMode => {
        console.log(debugMode);
        if (debugMode) document.oncontextmenu = showContextMenu; // Enable context menu if we're running in debug mode
    })

    Promise.all([
        window.backend.Bindings.CurrentVersion(),
        window.backend.Bindings.AvailableVersion()
    ]).then(result => {
        document.getElementById("CurrentVersion").innerText = result[0];
        document.getElementById("AvailableVersion").innerText = result[1];
    })
}

// We provide our entrypoint as a callback for runtime.Init
window.wails._.Init(start);