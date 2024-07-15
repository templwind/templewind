class Autocomplete extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Autocomplete connected");
	}
}

customElements.define("tw-autocomplete", Autocomplete);