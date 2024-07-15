class Pagination extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Pagination connected");
	}
}

customElements.define("tw-pagination", Pagination);