// class Testimoni {
//   quote = ""
//   image = ""

//   constructor(quote, image) {
//     this.quote = quote
//     this.image = image
//   }

//   get Quote() {
//     return this.quote
//   }

//   get Image() {
//     return this.image
//   }

//   get User() {
//     throw new Error("there is must be user to make testimoni")
//   }

//   get testimoniHTML() {
//     return `
//         <div class="card-testimoni">
//             <img src="${this.image}" class="profile-testimonial" />
//             <p class="quote">"${this.quote}"</p>
//             <p class="author">- ${this.user}</p>
//         </div>
//     `
//   }
// }

// class UserTestimoni extends Testimoni {
//   user = ""

//   constructor(user, quote, image) {
//     super(quote, image)
//     this.user = user
//   }

//   get User() {
//     return this.user
//   }
// }

// class CompanyTestimoni extends Testimoni {
//   company = ""

//   constructor(company, quote, image) {
//     super(quote, image)
//     this.company = company
//   }

//   get User() {
//     return "company: " + this.company
//   }
// }

const testimoniData = [
  {
    user: "lala",
    quote: "good",
    image:
      "https://images.unsplash.com/photo-1541562232579-512a21360020?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80",
    rating: 5,
  },
  {
    user: "lili",
    quote: " very good",
    image:
      "https://images.unsplash.com/photo-1689085383650-13d21072f5a6?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=774&q=80",
    rating: 3,
  },
  {
    user: "lulu",
    quote: "NOT Bad",
    image:
      "https://images.unsplash.com/photo-1688396068145-4bec246f615c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80",
    rating: 4,
  },
  {
    user: "lulu",
    quote: "NOT Bad",
    image:
      "https://images.unsplash.com/photo-1688396068145-4bec246f615c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80",
    rating: 2,
  },
]

const allTestimonial = () => {
  let testimoniHTML = ""

  testimoniData.forEach((card) => {
    testimoniHTML += `<div class="card-testimoni">
                        <img src="${card.image}" class="profile-testimonial" />
                        <p class="quote">"${card.quote}"</p>
                        <p class="author">- ${card.user}</p>
                      </div>`
  })
  document.getElementById("testimonial").innerHTML = testimoniHTML
}

allTestimonial()

const stars = (rating) => {
  let filterTestimoniHTML = ""

  const testimoniFilter = testimoniData.filter((card) => {
    return card.rating === rating
  })

  if (testimoniFilter.length === 0) {
    filterTestimoniHTML += `<h2>Data Not Found</h2>`
  } else {
    testimoniFilter.forEach((card) => {
      filterTestimoniHTML += `<div class="card-testimoni">
                        <img src="${card.image}" class="profile-testimonial" />
                        <p class="quote">"${card.quote}"</p>
                        <p class="author">- ${card.user}</p>
                      </div>`
    })
  }
  document.getElementById("testimonial").innerHTML = filterTestimoniHTML
}
