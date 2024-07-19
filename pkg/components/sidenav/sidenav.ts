import './sidenav.scss';

export class TwSidenav extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Sidenav connected");
	}
}

customElements.define("tw-sidenav", TwSidenav);