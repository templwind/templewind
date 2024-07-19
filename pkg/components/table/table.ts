import './table.scss';

export class TwTable extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Table connected");
	}
}

customElements.define("tw-table", TwTable);