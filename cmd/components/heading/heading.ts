class Heading extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Heading connected");
	}
}

customElements.define("tw-heading", Heading);