class AppAccordion extends HTMLElement {
    constructor() {
      super();
    }
  
    connectedCallback(): void {
      console.log("AppAccordion connected");
    }
  }
  
  customElements.define("app-accordion", AppAccordion);
  