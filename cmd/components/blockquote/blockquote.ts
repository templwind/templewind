class Blockquote extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Blockquote connected");
	}
}

customElements.define("tw-blockquote", Blockquote);