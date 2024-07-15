class Table extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Table connected");
	}
}

customElements.define("tw-table", Table);