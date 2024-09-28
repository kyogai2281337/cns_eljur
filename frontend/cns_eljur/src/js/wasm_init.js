import wasmURL from "/src/wasm/main.wasm?url";

const go = new Go();
WebAssembly.instantiateStreaming(fetch(wasmURL), go.importObject).then(result => {
    go.run(result.instance);
});

