const promise = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest()

  xhr.open("GET", "https://api.npoint.io/547dee21d86bed0412b9", true)
  xhr.onload = () => {
    if (xhr.status === 200) {
      resolve(JSON.parse(xhr.response))

    } else if (xhr.status === 400) {
      reject("Error loading Data")
    }
  }
  xhr.onerror = () => {
    reject("Network Error")
  }

  xhr.send()
})


let testimoniData = []

async function getData(rating) {
  try {
    const response = await promise
    testimoniData = response
    allTestimonial()
    console.log("respons", response)
  } catch (err) {
    console.log(err)
  }
}
getData()

function allTestimonial() {
  let testimoniHTML = ""

  testimoniData.forEach((card) => {
    testimoniHTML += `<div class="card" style="width: 18rem">
                          <img src="${card.image}" class="image-rounded" id="profile-img" />
                          <p class="quote text-end px-3">"${card.quote}"</p>
                          <p class="author text-end px-3">by ${card.user}</p>
                          <p class="ratings text-end px-3"><i class="fa-solid fa-star"></i> ${card.rating}</p>
                        </div>`
  })
  document.getElementById("testimonial").innerHTML = testimoniHTML
}

const stars = (rating) => {
  let filterTestimoniHTML = ""

  const testimoniFilter = testimoniData.filter((card) => {
    return card.rating === rating
  })

  if (testimoniFilter.length === 0) {
    filterTestimoniHTML += `<h2>Data Not Found</h2>`
  } else {
    testimoniFilter.forEach((card) => {
      filterTestimoniHTML += `<div class="card" style="width: 18rem">
                          <img src="${card.image}" class="image-rounded" id="profile-img" />
                          <p class="quote text-end px-3">"${card.quote}"</p>
                          <p class="author text-end px-3">- ${card.user}</p>
                          <p class="ratings text-end px-3"><i class="fa-solid fa-star"></i> ${card.rating}</p>
                        </div>`
    })
  }
  document.getElementById("testimonial").innerHTML = filterTestimoniHTML
}
