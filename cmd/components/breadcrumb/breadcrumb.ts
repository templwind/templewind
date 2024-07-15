class Breadcrumb extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Breadcrumb connected");
	}
}

customElements.define("tw-breadcrumb", Breadcrumb);