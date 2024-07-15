class List extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("List connected");
	}
}

customElements.define("tw-list", List);