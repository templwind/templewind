class Gallery extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Gallery connected");
	}
}

customElements.define("tw-gallery", Gallery);