class Toggle extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Toggle connected");
	}
}

customElements.define("tw-toggle", Toggle);