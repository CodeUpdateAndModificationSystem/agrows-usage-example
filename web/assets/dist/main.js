var u=Object.create;var{defineProperty:r,getPrototypeOf:v,getOwnPropertyNames:b}=Object;var g=Object.prototype.hasOwnProperty;var w=(e,t,n)=>{n=e!=null?u(v(e)):{};const o=t||!e||!e.__esModule?r(n,"default",{value:e,enumerable:!0}):n;for(let s of b(e))if(!g.call(o,s))r(o,s,{get:()=>e[s],enumerable:!0});return o};var f=(e,t)=>()=>(t||e((t={exports:{}}).exports,t),t.exports);var m=f((d,l)=>{var h=new MutationObserver((mutationsList)=>{function iterate(node){if(node.nodeType===Node.ELEMENT_NODE){const element=node;for(let i=0;i<element.attributes.length;i++){const attr=element.attributes[i];if(attr.name.startsWith("data-listen-")){const eventName=attr.name.replace("data-listen-",""),code=attr.value;console.warn(`Creating '${eventName}' event handler for '${element.tagName}' with code: ${code}`),element.addEventListener(eventName,(event)=>{eval(code)})}}element.childNodes.forEach(iterate)}}for(let e of mutationsList)e.addedNodes.forEach(iterate)});h.observe(document.body,{childList:!0,subtree:!0});console.warn("=== DOM observer started")});function E(e){a?.send(e)}var a,c=window;document.body.addEventListener("htmx:wsOpen",(e)=>{console.log("=== WebSocket opened"),a=e.detail.socketWrapper,c.ws=a,a?.send("reload?"),document.body.addEventListener("htmx:wsBeforeMessage",(n)=>{let o=n.detail.message;if(console.log("=== WebSocket message received: "+o),o==="reload!")console.log("=== Reloading page"),location.reload()}),a.onclose=()=>{console.warn("=== WebSocket closed"),c.onWsClose()}});window.sendMessage=E;var M=w(m(),1);async function W(){const e=new Go;let t;try{t=await WebAssembly.instantiateStreaming(fetch("/assets/main.wasm"),e.importObject),e.run(t.instance)}catch(n){console.error("WASM instantiation failed:",n)}}W().catch(console.error);console.log("=== Started");
