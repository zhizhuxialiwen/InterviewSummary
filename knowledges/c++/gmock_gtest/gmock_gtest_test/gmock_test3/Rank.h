#ifndef RANK_H_
#define RANK_H_

#include "IAPIProviderInterface.h"

namespace seamless {

class Rank {
public:
        virtual ~Rank() {}

public:
        void processQuery(IAPIProviderInterface* iAPIProvider);
};

}  // namespace seamless

#endif // RANK_H_