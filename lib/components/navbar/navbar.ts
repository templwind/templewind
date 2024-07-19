import './navbar.scss';

export class TwNavbar extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Navbar connected");
	}
}

customElements.define("tw-navbar", TwNavbar);