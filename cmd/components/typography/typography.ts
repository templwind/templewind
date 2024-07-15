class Typography extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Typography connected");
	}
}

customElements.define("tw-typography", Typography);