package comparator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func vPrintln(args ...interface{}) {
	if testing.Verbose() {
		fmt.Println(args...)
	}
}

// Takes a presser, and two clippings of an article based upon it:
// * Short, and representing "comment on PR" spam (artificially truncated from full clipping)
// * Full clipping representing real journalistic commentary
func TestBasicDupeDetection(t *testing.T) {
	C := NewMinhashComparator(0.80) // This is the threshold that detected the "comment on PR" clipping.
	assert.Equal(t, 250*8, len(C.Signature(doc1)))
	assert.NotEqual(t, C.Signature(doc1), C.Signature(doc2))
	vPrintln("Similarity of doc1|doc1:", C.Similarity(doc1, doc1))
	assert.True(t, C.Same(doc1, doc1))
	vPrintln("Similarity of doc1|doc2:", C.Similarity(doc1, doc2))
	assert.True(t, C.Same(doc1, doc2))
	vPrintln("Similarity of doc1|doc3:", C.Similarity(doc1, doc3))
	assert.False(t, C.Same(doc1, doc3))
	vPrintln("Similarity of doc2|doc3:", C.Similarity(doc2, doc3))
	assert.False(t, C.Same(doc2, doc3))
}

var (
	doc1 = `Already in the midst of the most significant expansion in the company’s celebrated history, Crystal Cruises’ next step in expanding its award-winning fleet is truly an historic endeavor. Together with the SS United States Conservancy, Crystal today announced it will save “America’s Flagship,” the SS United States, and embark on the enormous undertaking of bringing the ship into compliance with the latest standards, and returning her to oceangoing service. During the announcement, made at a press conference at the Manhattan Cruise Terminal in New York City, Crystal also committed to covering all costs associated with preserving the ship while undertaking a technical feasibility study, which is expected to be completed by the end of 2016.

  “The prospect of revitalizing the SS United States and reestablishing her as ‘America’s Flagship’ once again is a thrilling one. It will be a very challenging undertaking, but we are determined to apply the dedication and innovation that has always been the ship’s hallmark,” says Crystal President and CEO Edie Rodriguez. “We are honored to work with the SS United States Conservancy and government agencies in exploring the technical feasibility study so we can ultimately embark on the journey of transforming her into a sophisticated luxury cruise liner for the modern era.”`
	doc2 = `The following is the text of a news release from Crystal Cruises:

  (NEW YORK) — Already in the midst of the most significant expansion in the company’s celebrated history, Crystal Cruises’ next step in expanding its award-winning fleet is truly an historic endeavor. Together with the SS United States Conservancy, Crystal today announced it will save “America’s flagship,” the SS United States, and embark on the enormous undertaking of bringing the ship into compliance with the latest standards, and returning her to oceangoing service. During the announcement, made at a news conference at the Manhattan Cruise Terminal in New York City, Crystal also committed to covering all costs associated with preserving the ship while undertaking a technical feasibility study, which is expected to be completed by the end of 2016.

  “The prospect of revitalizing the SS United States and re-establishing her as ‘America’s flagship’ once again is a thrilling one. It will be a very challenging undertaking, but we are determined to apply the dedication and innovation that has always been the ship’s hallmark,” said Crystal President and CEO Edie Rodriguez. “We are honored to work with the SS United States Conservancy and government agencies in exploring the technical feasibility study so we can ultimately embark on the journey of transforming her into a sophisticated luxury cruise liner for the modern era.”

  “Crystal’s ambitious vision for the SS United States will ensure our nation’s flagship is once again a global ambassador for the highest standards of American innovation, quality and design,” said Susan Gibbs, executive director of the SS United States Conservancy and granddaughter of the ship’s designer, William Francis Gibbs.`
	doc3 = `The following is the text of a news release from Crystal Cruises:

  (NEW YORK) — Already in the midst of the most significant expansion in the company’s celebrated history, Crystal Cruises’ next step in expanding its award-winning fleet is truly an historic endeavor. Together with the SS United States Conservancy, Crystal today announced it will save “America’s flagship,” the SS United States, and embark on the enormous undertaking of bringing the ship into compliance with the latest standards, and returning her to oceangoing service. During the announcement, made at a news conference at the Manhattan Cruise Terminal in New York City, Crystal also committed to covering all costs associated with preserving the ship while undertaking a technical feasibility study, which is expected to be completed by the end of 2016.

  “The prospect of revitalizing the SS United States and re-establishing her as ‘America’s flagship’ once again is a thrilling one. It will be a very challenging undertaking, but we are determined to apply the dedication and innovation that has always been the ship’s hallmark,” said Crystal President and CEO Edie Rodriguez. “We are honored to work with the SS United States Conservancy and government agencies in exploring the technical feasibility study so we can ultimately embark on the journey of transforming her into a sophisticated luxury cruise liner for the modern era.”

  “Crystal’s ambitious vision for the SS United States will ensure our nation’s flagship is once again a global ambassador for the highest standards of American innovation, quality and design,” said Susan Gibbs, executive director of the SS United States Conservancy and granddaughter of the ship’s designer, William Francis Gibbs. We are thrilled that the SS United States is now poised to make a triumphant return to sea and that the ship’s historical legacy will continue to intrigue and inspire a new generation.”

  In order to meet modern demands and be in full regulatory compliance, the SS United States will have to be extensively rebuilt to meet over 60 years of new maritime rules and shipbuilding practices. The modern United States by Crystal Cruises will be transformed into an 800-guest-capacity vessel, featuring 400 luxurious suites measuring about 350 square feet with dining, entertainment, spa and other luxury guest amenities that are true to the ship’s storied history. Features of the original SS United States such as the Promenade and Navajo Lounge will be retained, while new engines and sophisticated marine technology will be installed to maintain its title as the fastest cruise vessel in the world.

  “It is truly a privilege for the world’s most awarded luxury cruise line to be entrusted with the opportunity of restoring a ship that served as a symbol of patriotism and maritime supremacy and bring her into the modern day, while also giving guests a taste of a bygone era of luxury travel,” Rodriguez said.

  Crystal will be examining exciting new itineraries for the 60,000-gross-ton United States by Crystal Cruises including not only the traditional trans-Atlantic voyages from New York City, but cruises from key U.S. ports as well as international voyages around the globe which are a signature offering of Crystal and part of the line’s “World Cruise.”

  The epitome of American post-war innovation and design, the SS United States was launched in 1952 and captured the trans-Atlantic speed record on its maiden voyage — a record to this day that still stands. It remains the largest passenger ship ever designed and built in America. Before its retirement in 1969, the SS United States was the most glamorous and elegant ship in the world, having transported four U.S. presidents, international royalty, many of Hollywood’s “golden era” celebrities, as well as a million passengers. While the ship captivated travelers with its features and elegance, the ship’s origin was equally intriguing. It was designed as part of a top-secret Pentagon program during the Cold War, which stipulated it could be quickly converted from a luxury liner into a naval troop ship in the event of a war, carrying 15,000 troops with a 240,000 shaft horsepower propulsion plant capable of traveling 10,000 nautical miles — almost halfway around the globe — without refueling.

  In October 2015, the SS United States Conservancy’s board of directors announced that the persistent challenge of covering the vessel’s monthly expenses had compelled them to engage a ship broker to explore the potential sale of the ship to be responsibly recycled. This news resulted in an outpouring of public support worldwide and led to the conservancy raising additional funds which enabled the organization to continue its preservation efforts and pursue negotiations with potential investors and partners.

  “The conservancy could never have reached this momentous milestone without the lifeline provided by our supporters from across the country and around the world. Thousands responded to our SOS last October and they refused to give up the fight for America’s flagship,” said Gibbs.

  To facilitate the complex technical feasibility study and to ensure a smooth execution of the project, Crystal has appointed retired U.S. Coast Guard Rear Adm. Tim Sullivan to build and lead a team with a wide range of cruise line technical, legal and regulatory expertise. With 36 years of active service, Sullivan has extensive experience in ship operations as a commanding officer of numerous Coast Guard cutters, and over the years has engaged in high level of interaction with a myriad of U.S. government agencies and international regulatory entities.

  “Tim’s integrity and leadership will help ensure the feasibility study is conducted with appropriately wide consultation, and rigorous adherence to both safety and environmental awareness,” said Rodriguez.

  The conservancy will continue to expand its curatorial and archival collections as it advances its mission of educating the public about the SS United States’ history. The organization will work with Crystal to establish shipboard displays and other educational programs. Planning is also underway for a land-based museum dedicated to preserving the legacy of America’s flagship along with broader design, innovation, and discovery themes. The museum will feature a wide range of original artifacts and historic components from the ship’s heyday.

  About SS United States Conservancy

  A national nonprofit organization, the SS United States Conservancy leads the global effort to save and repurpose America’s flagship, the SS United States. The conservancy raises public awareness and financial resources for the maintenance, restoration and ultimate reuse of this iconic vessel and works to ensure that the fastest ocean liner ever to cross the Atlantic remains an inspiration for generations to come. For more information about the SS United States, visit www.ssusc.org or the conservancy Facebook page.`
)
