class Button extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Button connected");
	}
}

customElements.define("tw-button", Button);