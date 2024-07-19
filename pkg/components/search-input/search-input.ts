import './search-input.scss';

export class TwSearchInput extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("SearchInput connected");
	}
}

customElements.define("tw-search-input", TwSearchInput);