import './mega-menu.scss';

export class TwMegaMenu extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("MegaMenu connected");
	}
}

customElements.define("tw-mega-menu", TwMegaMenu);