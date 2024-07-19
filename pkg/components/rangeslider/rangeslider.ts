import "./rangeslider.scss";

export class TwRangeSlider extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback(): void {
    console.log("RangeSlider connected");
  }
}

customElements.define("tw-range-slider", TwRangeSlider);
