class BottomNavigation extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("BottomNavigation connected");
	}
}

customElements.define("tw-bottom-navigation", BottomNavigation);