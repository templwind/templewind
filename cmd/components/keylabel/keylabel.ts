class Keylabel extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Keylabel connected");
	}
}

customElements.define("tw-keylabel", Keylabel);