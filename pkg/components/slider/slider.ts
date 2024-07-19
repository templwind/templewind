import './slider.scss';

export class TwSlider extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Slider connected");
	}
}

customElements.define("tw-slider", TwSlider);