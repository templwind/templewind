class AppHeader extends HTMLElement {
    constructor() {
      super();
    }
  
    connectedCallback(): void {
      console.log("AppHeader connected");
    }
  }
  
  customElements.define("app-header", AppHeader);
  