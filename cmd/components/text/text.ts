class Text extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Text connected");
	}
}

customElements.define("tw-text", Text);