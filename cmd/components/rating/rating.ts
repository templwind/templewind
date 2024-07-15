class Rating extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Rating connected");
	}
}

customElements.define("tw-rating", Rating);