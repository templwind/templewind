class SearchInput extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("SearchInput connected");
	}
}

customElements.define("tw-search-input", SearchInput);