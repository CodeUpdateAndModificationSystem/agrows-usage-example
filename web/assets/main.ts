import './wsManager'
import './domObserver'

async function loadWasm() {
    const go = new Go()
    let wasmModule
    
    try {
        wasmModule = await WebAssembly.instantiateStreaming(fetch("/assets/main.wasm"), go.importObject)
        go.run(wasmModule.instance)
    } catch (err) {
        console.error('WASM instantiation failed:', err)
    }
}

loadWasm().catch(console.error)

console.log("=== Started")
