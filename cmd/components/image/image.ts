class Image extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Image connected");
	}
}

customElements.define("tw-image", Image);