class InputField extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("InputField connected");
	}
}

customElements.define("tw-input-field", InputField);