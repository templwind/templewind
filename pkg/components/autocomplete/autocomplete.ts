import './autocomplete.scss';

export class TwAutocomplete extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Autocomplete connected");
	}
}

customElements.define("tw-autocomplete", TwAutocomplete);