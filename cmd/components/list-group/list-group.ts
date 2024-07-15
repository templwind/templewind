class ListGroup extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("ListGroup connected");
	}
}

customElements.define("tw-list-group", ListGroup);