import "./alert.scss";

export class TwAlert extends HTMLElement {
  private alert: HTMLElement | null = null;
  private closeBtn: HTMLElement | null = null;
  private hideTimeout: number = 3000; // Default to 3 seconds

  constructor() {
    super();
  }

  connectedCallback(): void {
    // console.log("TwAlert connected");
    this.alert = this.querySelector(".alert");

    // Get the hide duration from the attribute, if specified
    const hideDuration = this.getAttribute("hide-duration");
    if (hideDuration) {
      this.hideTimeout = parseInt(hideDuration, 10);
    }

    if (this.alert) {
      this.closeBtn = this.alert.querySelector(".close");
      if (this.closeBtn) {
        this.closeBtn.addEventListener("click", () => {
          if (this.alert) {
            this.alert.remove();
          }
        });
      }

      // Automatically hide the alert after the specified duration
      setTimeout(() => {
        if (this.alert) {
          this.alert.remove();
        }
      }, this.hideTimeout);
    }
  }
}

customElements.define("tw-alert", TwAlert);
