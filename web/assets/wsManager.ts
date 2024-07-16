let ws: WebSocket
const wasmInstance = window as any

document.body.addEventListener("htmx:wsOpen", (event: Event) => {
    console.log("=== WebSocket opened")

    const customEvent = event as CustomEvent
    ws = customEvent.detail.socketWrapper

    wasmInstance.ws = ws

    ws?.send("reload?")

    document.body.addEventListener("htmx:wsBeforeMessage", (event: any) => {
        let msg = event.detail.message as string
        console.log("=== WebSocket message received: " + msg)

        if (msg === "reload!") {
            console.log("=== Reloading page")
            location.reload()
        }
    })

    ws.onclose = () => {
        console.warn("=== WebSocket closed")
        wasmInstance.onWsClose()
    }
})

export function sendMessage(message: Uint8Array): void {
    ws?.send(message)
}

(window as any).sendMessage = sendMessage

export function getWebSocket(): WebSocket | null {
    return ws
}
