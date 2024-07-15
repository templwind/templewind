class Slider extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Slider connected");
	}
}

customElements.define("tw-slider", Slider);