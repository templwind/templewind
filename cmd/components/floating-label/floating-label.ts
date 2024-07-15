class FloatingLabel extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("FloatingLabel connected");
	}
}

customElements.define("tw-floating-label", FloatingLabel);