import './pagination.scss';

export class TwPagination extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Pagination connected");
	}
}

customElements.define("tw-pagination", TwPagination);