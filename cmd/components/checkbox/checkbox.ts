class Checkbox extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Checkbox connected");
	}
}

customElements.define("tw-checkbox", Checkbox);