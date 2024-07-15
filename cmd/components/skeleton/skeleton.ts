class Skeleton extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Skeleton connected");
	}
}

customElements.define("tw-skeleton", Skeleton);