class Testimoni {
  quote = ""
  image = ""

  constructor(quote, image) {
    this.quote = quote
    this.image = image
  }

  get Quote() {
    return this.quote
  }

  get Image() {
    return this.image
  }

  get User() {
    throw new Error("there is must be user to make testimoni")
  }

  get testimoniHTML() {
    return `
        <div class="card-testimoni">
            <img src="${this.image}" class="profile-testimonial" />
            <p class="quote">"${this.quote}"</p>
            <p class="author">- ${this.user}</p>
        </div>
    `
  }
}

class UserTestimoni extends Testimoni {
  user = ""

  constructor(user, quote, image) {
    super(quote, image)
    this.user = user
  }

  get User() {
    return this.user
  }
}

class CompanyTestimoni extends Testimoni {
  company = ""

  constructor(company, quote, image) {
    super(quote, image)
    this.company = company
  }

  get User() {
    return "company: " + this.company
  }
}

const testimoni1 = new UserTestimoni(
  "lala",
  "good job",
  "https://images.unsplash.com/photo-1541562232579-512a21360020?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80"
)
const testimoni2 = new UserTestimoni(
  "lala",
  "good job",
  "https://images.unsplash.com/photo-1689085383650-13d21072f5a6?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=774&q=80"
)
const testimoni3 = new CompanyTestimoni(
  "lala",
  "good job",
  "https://images.unsplash.com/photo-1688396068145-4bec246f615c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80"
)

let testimoniData = [testimoni1, testimoni2, testimoni3]

let testimoniHTML = ""

for (let i = 0; i < testimoniData.length; i++) {
  testimoniHTML += testimoniData[i].testimoniHTML
}
document.getElementById("testimonial").innerHTML = testimoniHTML
