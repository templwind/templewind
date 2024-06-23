class AppShell extends HTMLElement {
    constructor() {
      super();
    }
  
    connectedCallback(): void {
      console.log("AppShell connected");
    }
  }
  
  customElements.define("app-shell", AppShell);
  