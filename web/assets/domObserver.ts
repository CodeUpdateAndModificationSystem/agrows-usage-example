const observer = new MutationObserver((mutationsList: MutationRecord[]) => {
  function iterate(node: Node) {
    if (node.nodeType === Node.ELEMENT_NODE) {
      const element = node as Element;
      for (let i = 0; i < element.attributes.length; i++) {
        const attr = element.attributes[i];
        if (attr.name.startsWith("data-listen-")) {
          const eventName = attr.name.replace("data-listen-", "");
          const code = attr.value;
          console.warn(`Creating '${eventName}' event handler for '${element.tagName}' with code: ${code}`);
          element.addEventListener(eventName, (event: Event) => {
            eval(code);
          });
        }
      }
      element.childNodes.forEach(iterate);
    }
  }
  for (let mutation of mutationsList) {
    mutation.addedNodes.forEach(iterate);
  }
});

observer.observe(document.body, { childList: true, subtree: true });
console.warn("=== DOM observer started");
