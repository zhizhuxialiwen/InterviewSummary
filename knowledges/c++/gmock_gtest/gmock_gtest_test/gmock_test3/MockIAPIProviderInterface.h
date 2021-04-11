#ifndef MOCKIAPIPROVIDERINTERFACE_H_
#define MOCKIAPIPROVIDERINTERFACE_H_

#include <gmock/gmock.h>

#include "IAPIProviderInterface.h"
#include "IParameterInterface.h"

namespace seamless {

class MockIAPIProviderInterface: public IAPIProviderInterface{
public:
        MOCK_METHOD0(getParameterInterface, IParameterInterface*());
};

}  // namespace seamless

#endif // MOCKIAPIPROVIDERINTERFACE_H_