import './bottom-navigation.scss';

export class TwBottomNavigation extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("BottomNavigation connected");
	}
}

customElements.define("tw-bottom-navigation", TwBottomNavigation);