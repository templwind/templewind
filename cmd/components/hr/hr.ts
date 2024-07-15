class Hr extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Hr connected");
	}
}

customElements.define("tw-hr", Hr);