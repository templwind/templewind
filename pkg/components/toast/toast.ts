import './toast.scss';

export class TwToast extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Toast connected");
	}
}

customElements.define("tw-toast", TwToast);