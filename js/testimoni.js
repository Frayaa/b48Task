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
    rating: 4,
  },
]

const allTestimonial = () => {
  let testimoniHTML = ""

  testimoniData.forEach((card) => {
    testimoniHTML += `<div class="card-testimoni">
                        <img src="${card.image}" class="profile-testimonial" />
                        <p class="quote">"${card.quote}"</p>
                        <p class="author">by ${card.user}</p>
                        <p class="ratings"><i class="fa-solid fa-star"></i> ${card.rating}</p>
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
                        <p class="ratings">- ${card.rating}</p>
                      </div>`
    })
  }
  document.getElementById("testimonial").innerHTML = filterTestimoniHTML
}
