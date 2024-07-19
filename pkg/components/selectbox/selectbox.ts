import "./selectbox.scss";

export class TwSelectBox extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback(): void {
    console.log("Select connected");
  }
}

customElements.define("tw-select-box", TwSelectBox);
