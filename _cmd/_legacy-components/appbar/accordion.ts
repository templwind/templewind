class AppBar extends HTMLElement {
    constructor() {
      super();
    }
  
    connectedCallback(): void {
      console.log("AppBar connected");
    }
  }
  
  customElements.define("app-bar", AppBar);
  