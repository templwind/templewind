class Spinner extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Spinner connected");
	}
}

customElements.define("tw-spinner", Spinner);