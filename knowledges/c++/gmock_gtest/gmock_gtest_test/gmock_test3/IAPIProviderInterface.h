#ifndef IAPIPROVIDERINTERFACE_H_
#define IAPIPROVIDERINTERFACE_H_

#include <boost/cstdint.hpp>

#include "IParameterInterface.h"
#include "VariantField.h"

namespace seamless {

class IAPIProviderInterface {
public:
    IAPIProviderInterface() {}
    virtual ~IAPIProviderInterface() {}

public:
    virtual IParameterInterface* getParameterInterface() = 0;
};

}

#endif // IAPIPROVIDERINTERFACE_H_