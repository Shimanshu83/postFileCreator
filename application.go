package application

type Header struct {
	BatchId      string
	LengthRecord string
	DPID         string
}

type Details struct {
	LineNo     string
	DetailType string
	ClientId   string
	Addreshl1  string
	Addreshl2  string
	Pincode    string
	Email      string
	Password   string
}

/**
let create an easy first then we will think about anonmyous part this is
really important for me and all of you thanks you please check
Discusssion on how to create a datamapper
see we will be having datamapper which will alread have some predefined
data like In header most of the data except batch_number and no_of records will
be availaible.

In details most of the data except linenumber and other needs to be created at runtime from predefined data.

what will be the data source this also needs to be taken care of we need to
think about his particular thingh what will be the datasource and how to deal with that
stuff.

First we need to get all active data from some table.
then we need to creat a predefined database set which will be static and
check for fields and add it to our table.

**/

func AddressPositionalDataMapper() {

	// get some address record from here which is important once address record obtained
	// it should be in this manner
	// []Address Struct where Address is struct
	/**
		once we get address struct then we need to create header out of it
		which can be hardcoded for now we can create better way to solve this problem
		first do it easy way then we can approach this problem hard way too
		Header :-

		provide all the header data from

		header{
			BatchId batchIdGeneratorLogic
			DPID "value"
			LengthRecord len(Provided Data)
		}

		for index , value := range record {

			we need to get value from struct //
			we can also create a map for all the and only we need to do
			things which is important for me and all of us
			create a map first search for Name value in the struct
			replace it value with the existing details value and
			build that part iteratevly do for other parts too
			and keep doing that thingsh for other party too.
		}


	**/

}
