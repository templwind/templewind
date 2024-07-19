import './shell.scss';

export class TwShell extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Shell connected");
	}
}

customElements.define("tw-shell", TwShell);