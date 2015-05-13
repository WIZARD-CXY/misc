int Min(int *numbers, int len){
	if(numbers == null || len <=0){
		throw new std::exception("Invalid params")
	}
	if(len==1){
		return numbers[0];
	}

	int index1=0;
	int index2=len-1;

	int indexMid=index1;// set initial val here in case the numbers array is not rotated

	while(number[index1]>=number[index2]){
		if(index2-index1==1){
			indexMid=index2;
			break;
		}

		indexMid=(index1+index2)/2;

		if(numbers[indexMid]>=numbers[index1]){
			index1=indexMid;
		}else if(numbers[indexMid]<=numbers[index2]){
			index2=indexMid;
		}
	}
	return numbers[indexMid];
}