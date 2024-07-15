class Paragraph extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Paragraph connected");
	}
}

customElements.define("tw-paragraph", Paragraph);